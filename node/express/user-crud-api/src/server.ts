import "dotenv/config";

import app from "./app.js";
import prisma, { adapterAvailable } from "./config/database.js";

const PORT = Number(process.env.PORT) || 5000;

const startServer = async () => {
    try {
        if (adapterAvailable) {
            await prisma.$connect();
            console.log("Database connected successfully");
        } else {
            console.warn("@prisma/adapter-pg not installed — skipping database connection.");
        }

        app.listen(PORT, () => {
            console.log(
                `Server running on http://localhost:${PORT}`
            );
        });
    } catch (error) {
        console.error("Failed to start server:", error);

        if (adapterAvailable) await prisma.$disconnect();

        process.exit(1);
    }
};

startServer();