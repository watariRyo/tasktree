export const load = async ({ locals }) => {
	return {
		isLoggedIn: locals.user !== null
	};
};
