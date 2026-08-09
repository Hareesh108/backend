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
export declare const createUser: (data: CreateUserData) => Promise<{
    id: number;
    name: string;
    email: string;
    age: number | null;
    createdAt: Date;
    updatedAt: Date;
}>;
export declare const getAllUsers: () => Promise<{
    id: number;
    name: string;
    email: string;
    age: number | null;
    createdAt: Date;
    updatedAt: Date;
}[]>;
export declare const getUserById: (id: number) => Promise<{
    id: number;
    name: string;
    email: string;
    age: number | null;
    createdAt: Date;
    updatedAt: Date;
} | null>;
export declare const updateUser: (id: number, data: UpdateUserData) => Promise<{
    id: number;
    name: string;
    email: string;
    age: number | null;
    createdAt: Date;
    updatedAt: Date;
}>;
export declare const deleteUser: (id: number) => Promise<{
    id: number;
    name: string;
    email: string;
    age: number | null;
    createdAt: Date;
    updatedAt: Date;
}>;
//# sourceMappingURL=user.service.d.ts.map