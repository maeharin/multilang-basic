data class User(val id: Int, val name: String)

val users = listOf(
  User(0, "hoge"),
  User(1, "fuge"),
  User(2, "piyo")
)

val ids = users.map { it.id }
println(ids)

val names = users.map { it.name }
println(names)
