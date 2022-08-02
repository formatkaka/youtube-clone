CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  name VARCHAR(240) NOT NULL,
  email VARCHAR(100) NOT NULL,
  password VARCHAR(240) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
 );
 
 CREATE TABLE videos (
  id SERIAL PRIMARY KEY,
  title VARCHAR(240) NOT NULL,
  video_url VARCHAR(240) NOT NULL,
  user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
  views BIGINT DEFAULT 0,
  poster VARCHAR(240) NOT NULL,
  duration INTEGER NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
 );
  
 CREATE TABLE tags(
  id SERIAL PRIMARY KEY,
  name VARCHAR(40) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP 
 );
 
 CREATE TABLE video_tags_bridge(
  id SERIAL PRIMARY KEY,
  video_id INTEGER REFERENCES videos(id),
  tag_id INTEGER REFERENCES tags(id),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP  
 )