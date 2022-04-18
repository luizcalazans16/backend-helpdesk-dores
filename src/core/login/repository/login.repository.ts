import { LoginResponse } from "../dto/login.response";

export class LoginRepository {

    async attemptLogin(): Promise<{
        success: boolean;
        response?: LoginResponse;
        error?: Error;
    }> {
        try {
            const response = await
            return {
                success: true,
                response?: response,
            }
        } catch (e) {
            return {
                success: false,
                error?: e
            }
        }
    }
}