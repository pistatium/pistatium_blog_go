from google.cloud import datastore


def create_user():
    # gcloud auth application-default login
    client = datastore.Client()

    key = client.key("AdminUser", "kimihiro-n")
    au = datastore.Entity(key)
    au['password'] = '$2y$12$IXFmmeezqymra5O00L95lejHgVvMf9n7vDsKrU7f3s3zJTd5aePNS'
    client.put(au)
    print("Created user: kimihiro-n")


if __name__ == '__main__':
    create_user()
