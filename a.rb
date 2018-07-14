User = Struct.new(:id, :name)

users = [
  User.new(0, "hoge"),
  User.new(1, "fuge"),
  User.new(2, "piyo"),
]

ids = users.map(&:id)
p ids

names = users.map(&:name)
p names
