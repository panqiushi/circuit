export interface IAccordion {
	options?: {};
	show(): void;
	hide(): void;
}
export interface IBasePlugin<O, E> {
	el: E;
	options?: O;
	events?: {};
}
declare class HSBasePlugin<O, E = HTMLElement> implements IBasePlugin<O, E> {
	el: E;
	options: O;
	events?: any;
	constructor(el: E, options: O, events?: any);
	isIOS(): boolean;
	isIpadOS(): boolean;
	createCollection(collection: any[], element: any): void;
	fireEvent(evt: string, payload?: any): any;
	dispatch(evt: string, element: any, payload?: any): void;
	on(evt: string, cb: Function): void;
	afterTransition(el: HTMLElement, callback: Function): void;
	onTransitionEnd(el: HTMLElement, cb: Function): void;
	getClassProperty(el: HTMLElement, prop: string, val?: string): string;
	getClassPropertyAlt(el: HTMLElement, prop?: string, val?: string): string;
	htmlToElement(html: string): HTMLElement;
	classToClassList(classes: string, target: HTMLElement, splitter?: string): void;
	debounce(func: Function, timeout?: number): (...args: any[]) => void;
	checkIfFormElement(target: HTMLElement): boolean;
	static isEnoughSpace(el: HTMLElement, toggle: HTMLElement, preferredPosition?: "top" | "bottom" | "auto", space?: number, wrapper?: HTMLElement | null): boolean;
	static isParentOrElementHidden(element: any): any;
}
declare class HSAccordion extends HSBasePlugin<{}> implements IAccordion {
	private readonly toggle;
	content: HTMLElement | null;
	private readonly group;
	private readonly isAlwaysOpened;
	constructor(el: HTMLElement, options?: {}, events?: {});
	private init;
	show(): boolean;
	hide(): boolean;
	static getInstance(target: HTMLElement | string, isInstance?: boolean): HTMLElement | {
		id: number;
		element: HSAccordion;
	};
	static show(target: HTMLElement): void;
	static hide(target: HTMLElement): void;
	static autoInit(): void;
	static on(evt: string, target: HTMLElement, cb: Function): void;
}

export {
	HSAccordion as default,
};

export {};
