import { PrismaClient } from "../../generated/prisma/client.js";
import { createRequire } from "module";

const require = createRequire(import.meta.url);

let prismaInstance: any = null;
let initError: Error | null = null;

// Detect whether an adapter package is available synchronously.
let adapterAvailable = true;
let adapterName: string | null = null;
try {
    // Prefer `@prisma/adapter-ppg` if the user requested it, fall back to `@prisma/adapter-pg`.
    if (require.resolve("@prisma/adapter-ppg")) {
        adapterName = "@prisma/adapter-ppg";
    }
} catch (_) {
    try {
        if (require.resolve("@prisma/adapter-pg")) {
            adapterName = "@prisma/adapter-pg";
        }
    } catch (_e) {
        adapterName = null;
    }
}

if (!adapterName) adapterAvailable = false;

export { adapterAvailable };

function createPrismaInstance() {
    if (prismaInstance) return prismaInstance;

    const connectionString = process.env["DATABASE_URL"];

    if (!connectionString) {
        initError = new Error(
            "DATABASE_URL is not set. Set DATABASE_URL or pass `accelerateUrl` to PrismaClient."
        );
        prismaInstance = createPrismaStub(initError);
        return prismaInstance;
    }

    try {
        // Try to require the chosen adapter package.
        const mod = require(adapterName as string);

        // Prefer named export `PrismaPg`, then `PrismaPpg`, then default export.
        const PrismaCtor = mod.PrismaPg ?? mod.PrismaPpg ?? mod.default ?? mod;

        const adapter = new PrismaCtor({ connectionString });
        prismaInstance = new PrismaClient({ adapter });
        return prismaInstance;
    } catch (err: any) {
        initError = new Error(
            `${adapterName ?? "@prisma/adapter-pg"} is not installed or failed to load. Install it to enable DB connections.`
        );
        prismaInstance = createPrismaStub(initError);
        return prismaInstance;
    }
}

function createPrismaStub(err: Error) {
    return new Proxy(
        {},
        {
            get() {
                throw err;
            },
        }
    );
}

const prismaProxy = new Proxy(
    {},
    {
        get(_, prop) {
            const client = createPrismaInstance();
            const value = (client as any)[prop];
            // If the property is a function, bind it to the client so calls work.
            if (typeof value === "function") return value.bind(client);
            return value;
        },
        // Support `await prisma` checks in some libraries
        getOwnPropertyDescriptor() {
            createPrismaInstance();
            return Object.getOwnPropertyDescriptor(prismaInstance, "");
        },
    }
);

export default prismaProxy as any;