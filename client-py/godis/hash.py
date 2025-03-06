from .client import client

class HashStorage:
    def __init__(self, host: str = "localhost", port: int = 6480):
        self.host = host
        self.port = port

    def hget(self, key: str, field: str) -> str:
        """
        Run a command on the Godis server and return the response.
        """
        try:
            response = client(self.host, self.port, f"HGET {key} {field}")
            return response
        except Exception as e:
            return f"Error: {e}"
        
    def hset(self, key: str, **fields) -> str:
        """
        Sets one or more fields with values ​​to a Godis server hash key.
        """
        if not fields:
            return "Error: No fields provided for HSET."

        try:
            fields_str = " ".join(f"{field} {value}" for field, value in fields.items())

            command = f"HSET {key} {fields_str}"

            response = client(self.host, self.port, command)

            return response
        
        except Exception as e:
            return f"Error: {e}"

