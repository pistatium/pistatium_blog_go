from google.cloud import datastore



def create_conf():
    # gcloud auth application-default login
    client = datastore.Client(project='local-app')

    query = client.query(kind='Blog')
    query_iterator = query.fetch(limit=10)
    for q in query_iterator:
        print(q)

    key = client.key("Conf", 1)
    c = datastore.Entity(key)
    c['secret'] = 'P5Xau2@v6U8mv4wiT6Y9Evj%R8aumgR8'
    client.put(c)
    print("Created conf", from_datastore(c))


def from_datastore(entity):

    if not entity:
        return None
    entity['id'] = entity.key.id
    return entity


def create_user():
    # gcloud auth application-default login
    client = datastore.Client()

    key = client.key("AdminUser", "kimihiro-n")
    au = datastore.Entity(key)
    au['password'] = '$2y$12$IXFmmeezqymra5O00L95lejHgVvMf9n7vDsKrU7f3s3zJTd5aePNS'
    client.put(au)
    print("Created user: kimihiro-n")


if __name__ == '__main__':
    #create_user()
    create_conf()
