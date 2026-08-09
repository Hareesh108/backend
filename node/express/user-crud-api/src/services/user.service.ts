import prisma from "../config/database.js";

export interface CreateUserData {
    name: string;
    email: string;
    age?: number;
}

export interface UpdateUserData {
    name?: string;
    email?: string;
    age?: number;
}

export const createUser = async (data: CreateUserData) => {
    try {
        console.log('user.service.createUser: data=', data);
        const result = await prisma.user.create({ data });
        console.log('user.service.createUser: result=', result);
        return result;
    } catch (err) {
        console.error('user.service.createUser: error=', err);
        throw err;
    }
};

export const getAllUsers = async () => {
    return prisma.user.findMany({
        orderBy: {
            id: "desc"
        }
    });
};

export const getUserById = async (id: number) => {
    return prisma.user.findUnique({
        where: {
            id
        }
    });
};

export const updateUser = async (
    id: number,
    data: UpdateUserData
) => {
    return prisma.user.update({
        where: {
            id
        },
        data
    });
};

export const deleteUser = async (id: number) => {
    return prisma.user.delete({
        where: {
            id
        }
    });
};