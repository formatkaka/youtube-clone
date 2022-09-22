
import type { GetSession, Handle } from '@sveltejs/kit'
import * as cookie from 'cookie';

export const handle: Handle = async ({ event, resolve }: any) => {
	const cookieHeader = event.request.headers.get('cookie');
	const cookies = cookie.parse(cookieHeader ?? '');

	console.log('🚀 ~ file: hooks.server.ts ~ line 12 ~ cookies.token', cookies.token);
	if (!cookies.token) {
		return await resolve(event);
	}

	if (cookies.token) {
		event.locals.token = { username: cookies.token };
	}

	return await resolve(event);
};

export const getSession: GetSession = ({ locals }) => {
	if (!locals.token) return {}
  
	return {
	  user: {
		username: locals.token.username,
	  },
	}
  }