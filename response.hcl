encryption_key = ""

postgresql {
  host = env("DATABASE_HOST")
  port = env("DATABASE_PORT")
}