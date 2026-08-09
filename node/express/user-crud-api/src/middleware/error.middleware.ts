import type {
    Request,
    Response,
    NextFunction
} from "express";

export const errorHandler = (
    error: unknown,
    _req: Request,
    res: Response,
    _next: NextFunction
) => {
    // Log detailed error information
    if (error instanceof Error) {
        console.error('Error message:', error.message);
        console.error('Error stack:', error.stack);
    } else {
        console.error('Unknown error:', error);
    }

    const isDev = process.env.NODE_ENV !== 'production';
    return res.status(500).json({
        success: false,
        message: isDev && error instanceof Error ? error.message : 'Internal server error',
        ...(isDev && error instanceof Error ? { stack: error.stack } : {})
    });
};