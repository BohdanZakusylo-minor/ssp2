import pika
import sys

def log(message):
    """Print with prefix for better visibility in Docker logs"""
    print(f"[PYTHON-CONSUMER] {message}", flush=True)

def main_func():
    log("\n=== South Park Messages ===")

    try:
        credentials = pika.PlainCredentials("user", "password")
        params = pika.ConnectionParameters(
            host="rabbit-mq",          
            port=5672,
            virtual_host="/",         
            credentials=credentials,
            heartbeat=30,
            blocked_connection_timeout=30,
        )

        connection = pika.BlockingConnection(params)
        channel = connection.channel()

        channel.queue_declare(queue="southpark_messages", durable=True)

        def callback(ch, method, properties, body):
            log(f" [x] Received: {body.decode()}")

        channel.basic_consume(
            queue="southpark_messages",
            on_message_callback=callback,
            auto_ack=True,   
        )

        log("--Connected to Rabbitmq")
        log("--Waiting for messages")
        channel.start_consuming()

    except pika.exceptions.AMQPConnectionError as e:
        log(f"Connection error: {e}")
        sys.exit(1)

    except KeyboardInterrupt:
        log("\nInterrupted by user")
        if 'channel' in locals():
            channel.stop_consuming()
        if 'connection' in locals():
            connection.close()
        sys.exit(0)

    except Exception as e:
        log(f"Unexpected error: {e}")
        sys.exit(1)
