import prisma from "../config/database.js";
export const createUser = async (data) => {
    return prisma.user.create({
        data
    });
};
export const getAllUsers = async () => {
    return prisma.user.findMany({
        orderBy: {
            id: "desc"
        }
    });
};
export const getUserById = async (id) => {
    return prisma.user.findUnique({
        where: {
            id
        }
    });
};
export const updateUser = async (id, data) => {
    return prisma.user.update({
        where: {
            id
        },
        data
    });
};
export const deleteUser = async (id) => {
    return prisma.user.delete({
        where: {
            id
        }
    });
};
//# sourceMappingURL=user.service.js.map