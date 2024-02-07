import { ApiError, type ErrorResponseSchema } from '../../models/api.error';

export const responseAnalysis = async <T>(response: Response): Promise<T> => {
	// const headers = {
	// 	'content-type': 'application/json'
	// };

	// const response = await fetch(endPoint, {
	// 	mode: 'cors',
	// 	credentials: 'include', // サイトを跨ぐCookieの保持に必須
	// 	...config,
	// 	headers: {
	// 		...headers,
	// 		...config?.headers
	// 	}
	// }).catch(() => {
	// 	return Promise.reject();
	// });

	const contentType = response.headers.get('Content-Type') || '';

	if (!response.ok) {
		console.log('ng');
		const serverErrorContent = isJson(contentType)
			? ((await response.json()) as ErrorResponseSchema)
			: undefined;
		return Promise.reject(new ApiError(response, serverErrorContent));
	}

	if (isJson(contentType)) return await response.json();

	return Promise.resolve() as unknown as Promise<T>;
};

const isJson = (contentType: string) => contentType.includes('application/json');
