import { redirect } from '@sveltejs/kit';
import type { PageLoad } from './$types';

export const trailingSlash = 'ignore';

export const load: PageLoad = async ({ url }) => {
	throw redirect(308, `127.0.0.1://oauth2/?${url.searchParams.toString()}`);
};
