import yaml

with open("app/config/config.yaml", "r") as file:
    config = yaml.safe_load(file)

with open(".env", "w") as env_file:
    env_file.write(f"SERVER_PORT={config['server']['port']}\n")
    env_file.write(f"POSTGRES_HOST={config['database']['postgres']['host']}\n")
    env_file.write(f"POSTGRES_PORT={config['database']['postgres']['port']}\n")
    env_file.write(f"POSTGRES_USER={config['database']['postgres']['user']}\n")
    env_file.write(f"POSTGRES_PASSWORD={config['database']['postgres']['password']}\n")
    env_file.write(f"POSTGRES_DBNAME={config['database']['postgres']['dbname']}\n")
    env_file.write(f"POSTGRES_SSLMODE={config['database']['postgres']['sslmode']}\n")
    env_file.write(f"MONGO_URI={config['database']['mongo']['uri']}\n")
    env_file.write(f"MONGO_DATABASE={config['database']['mongo']['database']}\n")
    env_file.write(f"REDIS_HOST={config['redis']['host']}\n")
    env_file.write(f"REDIS_PORT={config['redis']['port']}\n")
    env_file.write(f"REDIS_PASSWORD={config['redis']['password']}\n")
    env_file.write(f"REDIS_DB={config['redis']['db']}\n")
    env_file.write(f"LOG_DIR={config['logging']['log_dir']}\n")
