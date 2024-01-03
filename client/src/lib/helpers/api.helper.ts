import { ApiError, type ErrorResponseSchema } from '../../models/api.error';

export const client = async <T>(endPoint: RequestInfo, config?: RequestInit): Promise<T> => {
	const headers = {
		'content-type': 'application/json'
	};

	const response = await fetch(endPoint, {
		mode: 'cors',
		credentials: 'include',
		...config,
		headers: {
			...headers,
			...config?.headers
		}
	}).catch(() => {
		return Promise.reject();
	});

	const contentType = response.headers.get('Content-Type') || '';

	if (!response.ok) {
		const serverErrorContent = isJson(contentType)
			? ((await response.json()) as ErrorResponseSchema)
			: undefined;
		return Promise.reject(new ApiError(response, serverErrorContent));
	}

	if (isJson(contentType)) return await response.json();

	return Promise.resolve() as unknown as Promise<T>;
};

const isJson = (contentType: string) => contentType.includes('application/json');
