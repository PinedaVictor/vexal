
# Main function
import sys
import os
from typing import Union
from fastapi import FastAPI
import routes.user

from dotenv import load_dotenv
load_dotenv()  # take environment variables from .env.

from skyflow.vault import Client, Configuration, RedactionType
from skyflow.errors import SkyflowError
from key import token_provider

app = FastAPI()
user = routes.user.router
app.include_router(user)


def env_check():
    if (sys.prefix != sys.base_prefix):
        print("Env initialized")
        return
    print("Env not initialized")
    return

@app.get("/")
def read_root():
    env_check()
    token_provider()
    return {"Hello": "World"}


@app.get("/items/{item_id}")
def read_item(item_id: int, q: Union[str, None] = None):
    return {"item_id": item_id, "q": q}


# Sample user id: 756f576f-1712-45c2-96aa-18e83ef743e7
@app.get("/user")
# NOTE: Id can be an array
def get_user(id: str = "0"):
    data = {}
    print("getting user id:", id)
    try:
        print("Confuiguring client")
        config = Configuration(
            os.environ["VAULT_ID"], os.environ["VAULT_URL"], token_provider)
        client = Client(config)

        data = {"records": [
            {
                "ids": [id],
                "table": "persons",
                "redaction": RedactionType.DEFAULT
            }
        ]}
        
        print("calling client.get")
        response = client.get(data)
        print('Response:', response["records"])
        data = response["records"]
    except SkyflowError as e:
        print('Error Occurred:',e, "----data:", e.data)
        if len(e.data["records"]) == 0:
            data = {"message:" "record not found"}
            return
        data = e.data
        
    return {"user": data}