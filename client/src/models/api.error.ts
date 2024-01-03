const API_ERROR = 'API_ERROR';

export type ErrorResponseSchema = {
	stauts: number;
	code?: string;
	message: string;
};

export class ApiError extends Error {
	name: string;
	url?: string;
	status?: number;
	statusText?: string;
	serverErrorContent?: ErrorResponseSchema;
	constructor(response?: Response, serverErrorContent?: ErrorResponseSchema) {
		super(response?.statusText || 'network error');
		this.name = API_ERROR;
		this.status = response?.status;
		this.statusText = response?.statusText;
		this.url = response?.url;
		this.serverErrorContent = serverErrorContent;
	}
	serialize() {
		return Object.assign({}, this);
	}
}
