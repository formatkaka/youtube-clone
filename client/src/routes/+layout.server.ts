import * as cookie from 'cookie';
import { getUserDetails } from '@services/auth';

export async function load({request}: any) {
    const cookieHeader = request.headers.get('cookie');
	const cookies = cookie.parse(cookieHeader ?? '');

    if (!cookies.token) {
		return {};
	}

    const userData = await getUserDetails(cookies.token)
    return {
        user: userData
    }
}