import * as userService from "../services/user.service.js";
export const createUser = async (req, res, next) => {
    try {
        const { name, email, age } = req.body;
        if (!name || !email) {
            return res.status(400).json({
                success: false,
                message: "Name and email are required"
            });
        }
        const user = await userService.createUser({
            name,
            email,
            age
        });
        return res.status(201).json({
            success: true,
            message: "User created successfully",
            data: user
        });
    }
    catch (error) {
        next(error);
    }
};
export const getAllUsers = async (_req, res, next) => {
    try {
        const users = await userService.getAllUsers();
        return res.status(200).json({
            success: true,
            data: users
        });
    }
    catch (error) {
        next(error);
    }
};
export const getUserById = async (req, res, next) => {
    try {
        const id = Number(req.params.id);
        if (Number.isNaN(id)) {
            return res.status(400).json({
                success: false,
                message: "Invalid user ID"
            });
        }
        const user = await userService.getUserById(id);
        if (!user) {
            return res.status(404).json({
                success: false,
                message: "User not found"
            });
        }
        return res.status(200).json({
            success: true,
            data: user
        });
    }
    catch (error) {
        next(error);
    }
};
export const updateUser = async (req, res, next) => {
    try {
        const id = Number(req.params.id);
        if (Number.isNaN(id)) {
            return res.status(400).json({
                success: false,
                message: "Invalid user ID"
            });
        }
        const { name, email, age } = req.body;
        const existingUser = await userService.getUserById(id);
        if (!existingUser) {
            return res.status(404).json({
                success: false,
                message: "User not found"
            });
        }
        const user = await userService.updateUser(id, {
            name,
            email,
            age
        });
        return res.status(200).json({
            success: true,
            message: "User updated successfully",
            data: user
        });
    }
    catch (error) {
        next(error);
    }
};
export const deleteUser = async (req, res, next) => {
    try {
        const id = Number(req.params.id);
        if (Number.isNaN(id)) {
            return res.status(400).json({
                success: false,
                message: "Invalid user ID"
            });
        }
        const existingUser = await userService.getUserById(id);
        if (!existingUser) {
            return res.status(404).json({
                success: false,
                message: "User not found"
            });
        }
        await userService.deleteUser(id);
        return res.status(200).json({
            success: true,
            message: "User deleted successfully"
        });
    }
    catch (error) {
        next(error);
    }
};
//# sourceMappingURL=user.controller.js.map