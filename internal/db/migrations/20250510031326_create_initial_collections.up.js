db.createCollection("eshops");
db.createCollection("merchants");
db.createCollection("likes");
db.createCollection("reports");

// Optional: add indexes
db.eshops.createIndex({ name: 1 }, { unique: true });
db.likes.createIndex({ userId: 1, merchantId: 1 }, { unique: true });
