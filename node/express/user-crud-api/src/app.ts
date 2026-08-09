import express from "express";
import cors from "cors";

import userRoutes from "./routes/user.routes.js";
import { errorHandler } from "./middleware/error.middleware.js";

const app = express();

app.use(cors());

app.use(express.json());

app.get("/", (_req, res) => {
    res.status(200).json({
        success: true,
        message: "User CRUD API is running"
    });
});

app.use("/api/users", userRoutes);

app.use(errorHandler);

export default app;