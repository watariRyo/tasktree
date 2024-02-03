export const load = async ({ locals }) => {
	if (locals.user !== null) {
		return {
			data: 'hoge2'
		};
	} else {
		return null;
	}
};
