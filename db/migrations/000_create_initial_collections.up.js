// 000_create_initial_collections.up.js

if (!db.getCollectionNames().includes("eshops")) {
  db.createCollection("eshops");
}

if (!db.getCollectionNames().includes("merchants")) {
  db.createCollection("merchants");
}

if (!db.getCollectionNames().includes("likes")) {
  db.createCollection("likes");
}

if (!db.getCollectionNames().includes("reports")) {
  db.createCollection("reports");
}
