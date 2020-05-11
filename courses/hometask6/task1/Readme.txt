Web server handling any request to it’s address and returning JSON like:
{"host": "127.0.0.1:8080", "user_agent": "curl/7.67.0", "request_uri": "/anything/you?want", "headers": {"Accept": ["*/*"], "User-Agent": ["curl/7.67.0"] } }

All data must be taken from request struct.
