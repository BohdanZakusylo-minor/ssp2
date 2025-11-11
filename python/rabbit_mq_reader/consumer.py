import pika
import sys

def main_func():
    print("\n=== South Park Messages ===")

    try:
        credentials = pika.PlainCredentials("user", "password")
        params = pika.ConnectionParameters(
            host="localhost",          
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
            print(f" [x] Received: {body.decode()}")

        channel.basic_consume(
            queue="southpark_messages",
            on_message_callback=callback,
            auto_ack=True,   
        )

        print("--Connected to Rabbitmq")
        print("--Waiting for messages")
        channel.start_consuming()

    except pika.exceptions.AMQPConnectionError as e:
        print(f"Connection error: {e}")
        sys.exit(1)

    except KeyboardInterrupt:
        print("\nInterrupted by user")
        if 'channel' in locals():
            channel.stop_consuming()
        if 'connection' in locals():
            connection.close()
        sys.exit(0)

    except Exception as e:
        print(f"Unexpected error: {e}")
        sys.exit(1)

