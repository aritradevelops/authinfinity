data "external_schema" "gorm" {
  program = [
    "go",
    "run",
    "-mod=mod",
    "./cmd/atlas-provider",
  ]
}
env "gorm" {
  src = data.external_schema.gorm.url
  dev = "docker://postgres/15/dev?search_path=public"
  migration {
    dir = "file://migrations"
  }
  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}
env "dev" {
  url = "postgresql://postgres:admin@localhost:5432/authinfinity?sslmode=disable"
  migration {
    dir = "file://migrations"
  }
}
env "local" {
  url = "postgresql://postgres:admin@localhost:5432/authinfinity?sslmode=disable"
  migration {
    dir = "file://migrations"
  }
}