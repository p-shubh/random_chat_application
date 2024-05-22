# Real-Time Chat Application Backend with Go

## Setup Instructions

1. **Clone the repository**:
    ```sh
    git clone https://github.com/p-shubh/chat-app.git
    cd chat-app
    ```

2. **Set up Supabase**:
    - Create a new project on Supabase.
    - Create the necessary tables for users and messages.
    - Obtain your Supabase DSN and set it as an environment variable.

3. **Configure environment variables**:
    ```sh
    export SUPABASE_DSN="your_supabase_dsn"
    ```

4. **Run setup script**:
    ```sh
    ./scripts/setup.sh
    ```

5. **Run the application**:
    ```sh
    go run cmd/app/main.go
    ```

## Usage

- Connect to the WebSocket endpoint at `ws://localhost:8080/ws` to start chatting.
- The backend will handle user pairing and message relaying in real-time.

## Notes

- Ensure your frontend (Next.js) is properly configured to connect to the WebSocket endpoint.
- The application is designed to be scalable and can handle multiple concurrent connections.
