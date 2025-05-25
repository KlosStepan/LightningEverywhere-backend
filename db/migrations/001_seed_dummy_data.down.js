// 001_seed_dummy_data.down.js

module.exports.down = async function (db) {
  // Delete inserted e-shops by their ids
  await db.collection('eshops').deleteMany({
    id: {
      $in: [
        "54ae38db-2a48-42a9-a569-e9e48ea90d46",
        "e20d248d-69cc-482d-9ce1-0a15cbe3bc2a",
        "b25efd17-3a50-4513-9dad-af47bbcfb687"
      ]
    }
  });

  // Delete inserted merchants by their properties.id
  await db.collection('merchants').deleteMany({
    "properties.id": {
      $in: [
        "63977929-fc0e-4695-9a04-3156e9d24c54",
        "10417593-72da-4ba3-ace0-93f6ba676b9e"
      ]
    }
  });
};
