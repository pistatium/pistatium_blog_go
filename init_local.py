import os
from google.cloud import datastore
import google
from google.auth.credentials import AnonymousCredentials

"""
export DATASTORE_EMULATOR_HOST=0.0.0.0:8059
export PROJECT_ID=local-app
"""

def create_conf(client):
    # gcloud auth application-default login
    key = client.key("Conf", '1')
    c = datastore.Entity(key)
    c['secret'] = 'P5Xau2@v6U8mv4wiT6Y9Evj%R8aumgR8'
    client.put(c)
    print("Created conf", from_datastore(c))


def from_datastore(entity):
    if not entity:
        return None
    entity['id'] = entity.key.id
    return entity


def create_user(client):
    # gcloud auth application-default login
    key = client.key("AdminUser", "kimihiro-n")
    au = datastore.Entity(key)
    au['password'] = '$2y$12$IXFmmeezqymra5O00L95lejHgVvMf9n7vDsKrU7f3s3zJTd5aePNS'
    client.put(au)
    print("Created user: kimihiro-n")


if __name__ == '__main__':
    client = datastore.Client(
        _http=None,
        project=os.environ.get('DATASTORE_PROJECT_ID')
    )

    create_user(client)
    create_conf(client)
