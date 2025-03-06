import socket


def client(host, port, command):
    with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as client_socket:
        client_socket.connect((host, port))

        client_socket.sendall(f"{command}\n".encode("utf-8"))

        response = client_socket.recv(1024).decode("utf-8")

        return response
