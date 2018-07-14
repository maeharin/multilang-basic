class User():
    def __init__(self, id, name):
        self.id = id
        self.name = name

users = [
  User(0, "hoge"),
  User(1, "fuge"),
  User(2, "piyo"),
]

ids = [user.id for user in users]
print(ids)

names = [user.name for user in users]
print(names)
