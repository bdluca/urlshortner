db.createUser(
    {
        user: "bluca",
        pwd: "bluca",
        roles: [
            {
                role: "readWrite",
                db: "test"
            }
        ]
    }
);