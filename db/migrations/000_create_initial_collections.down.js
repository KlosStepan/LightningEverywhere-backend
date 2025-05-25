// 000_create_initial_collections.down.js

/*if (db.getCollectionNames().includes("reports")) {
  db.reports.drop();
}

if (db.getCollectionNames().includes("likes")) {
  db.likes.drop();
}

if (db.getCollectionNames().includes("merchants")) {
  db.merchants.drop();
}

if (db.getCollectionNames().includes("eshops")) {
  db.eshops.drop();
}*/

db.eshops.drop();
db.merchants.drop();
db.likes.drop();
db.reports.drop();
