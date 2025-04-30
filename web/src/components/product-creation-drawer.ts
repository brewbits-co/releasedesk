import {html, LitElement} from 'lit';
import {customElement, state} from 'lit/decorators.js';
import {convertToSlug, encodeFormData} from "../utils";
import {createRef, Ref, ref} from 'lit/directives/ref.js';
import {map} from 'lit/directives/map.js';

@customElement('product-creation-drawer')
export class ProductCreationDrawer extends LitElement {
    override createRenderRoot() {
        return this;
    }

    @state()
    submitFailed = false;

    errorAlertRef: Ref<HTMLDivElement> = createRef();

    @state()
    errorData: ErrorResponse = {message: '', helpTexts: []};

    handleSubmit(event: SubmitEvent) {
        event.preventDefault();
        this.submitFailed = false;

        const formData = new FormData(event.target as HTMLFormElement);

        fetch('/internal/products', {
            method: 'POST',
            body: encodeFormData(formData),
            credentials: 'include',
            headers: {
                'Content-Type': 'application/x-www-form-urlencoded',
            },
        }).then(async response => {
            if (!response.ok) {
                this.submitFailed = true;
                this.errorData = await response.json();
                // wait form render to complete
                await this.updateComplete;
                // scroll the page to the error alert
                this.errorAlertRef.value?.scrollIntoView();
            } else {
                window.location.href = `/dashboard/${formData.get("Slug")}`;
            }
        });
    }

    handleSlugChange(e: any) {
        const inputValue = e.target.value;
        e.target.value = convertToSlug(inputValue); // Update the slug property
    }

    handleOpen() {
        // @ts-ignore
        window.productCreationDrawer.showModal()
    }

    handleClose() {
        // @ts-ignore
        window.productCreationDrawer.close()
        this.submitFailed = false;
    }

    override render() {
        // @ts-ignore
        return html`
            <div class="text-center">
                <button type="button"
                        class="inline-flex items-center rounded-md bg-violet-900 px-3 py-2 text-sm font-semibold text-white shadow-sm hover:bg-violet-800 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-violet-900"
                        @click="${() => this.handleOpen()}">
                    <i class="fa-solid fa-layer-group"></i>&nbsp;&nbsp;Create Your Product
                </button>
            </div>
            
            <dialog-drawer ref="productCreationDrawer" title="ok" description="yo"></dialog-drawer>

            <dialog id="productCredationDrawer"
                    class="relative z-10 opacity-0 transition-opacity duration-300 ease-in-out open:opacity-100"
                    aria-labelledby="slide-over-title" role="dialog"
                    aria-modal="true">
                <!-- Background backdrop, show/hide based on slide-over state. -->
                <div class="fixed inset-0 bg-gray-500/75"></div>

                <div class="fixed inset-0 overflow-hidden">
                    <div class="absolute inset-0 overflow-hidden">
                        <div class="pointer-events-none fixed inset-y-0 right-0 flex max-w-full pl-10">
                            <!--
                              Slide-over panel, show/hide based on slide-over state.
                    
                              Entering: "transform transition ease-in-out duration-500 sm:duration-700"
                                From: "translate-x-full"
                                To: "translate-x-0"
                              Leaving: "transform transition ease-in-out duration-500 sm:duration-700"
                                From: "translate-x-0"
                                To: "translate-x-full"
                            -->
                            <div class="pointer-events-auto w-screen max-w-md">
                                <div class="flex h-full flex-col overflow-y-auto bg-white shadow-xl">
                                    <div class="bg-violet-900 px-4 py-6 sm:px-6">
                                        <div class="flex items-center justify-between">
                                            <h2 class="text-base font-semibold text-white" id="slide-over-title">New
                                                Product</h2>
                                            <div class="ml-3 flex h-7 items-center">
                                                <button type="button"
                                                        @click="${() => this.handleClose()}"
                                                        class="relative rounded-md bg-violet-900 text-violet-200 hover:text-white focus:outline-none focus:ring-2 focus:ring-white">
                                                    <span class="absolute -inset-2.5"></span>
                                                    <span class="sr-only">Close panel</span>
                                                    <svg class="size-6" fill="none" viewBox="0 0 24 24"
                                                         stroke-width="1.5" stroke="currentColor" aria-hidden="true"
                                                         data-slot="icon">
                                                        <path stroke-linecap="round" stroke-linejoin="round"
                                                              d="M6 18 18 6M6 6l12 12"/>
                                                    </svg>
                                                </button>
                                            </div>
                                        </div>
                                        <div class="mt-1">
                                            <p class="text-sm text-violet-300">Get started by filling in the information
                                                below to create your new product.</p>
                                        </div>
                                    </div>
                                    <div class="relative flex-1 px-4 sm:px-6">
                                        <form slot="body" id="product-creation-form"
                                              @reset="${this.handleClose}"
                                              @submit="${this.handleSubmit}">
                                            <div class="space-y-6 pb-5 pt-6">
                                                <div>
                                                    <label for="project-name"
                                                           class="block text-sm/6 font-medium text-gray-900">Name</label>
                                                    <div class="mt-2">
                                                        <input type="text"
                                                               name="Name"
                                                               id="Name"
                                                               class="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-violet-800 sm:text-sm/6"
                                                               maxlength="255"
                                                               autocomplete="off"
                                                               required>
                                                    </div>
                                                </div>
                                                <div>
                                                    <label for="project-name"
                                                           class="block text-sm/6 font-medium text-gray-900">Slug</label>
                                                    <div class="mt-2">
                                                        <input type="text"
                                                               name="Slug"
                                                               id="Slug"
                                                               @input="${this.handleSlugChange}"
                                                               class="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-violet-800 sm:text-sm/6"
                                                               maxlength="255"
                                                               autocomplete="off"
                                                               required>
                                                    </div>
                                                </div>
                                                <div>
                                                    <label for="description"
                                                           class="block text-sm/6 font-medium text-gray-900">Description</label>
                                                    <div class="mt-2">
                <textarea id="Description"
                          name="Description"
                          rows="4"
                          autocomplete="off"
                          class="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-violet-800 sm:text-sm/6"></textarea>
                                                    </div>
                                                </div>
                                                <fieldset>
                                                    <legend class="text-sm/6 font-medium text-gray-900">Visibility
                                                    </legend>
                                                    <div class="mt-2 space-y-4">
                                                        <div class="relative flex items-start">
                                                            <div class="absolute flex h-6 items-center">
                                                                <input id="Private"
                                                                       name="Private"
                                                                       value="false"
                                                                       aria-describedby="privacy-public-description"
                                                                       type="radio"
                                                                       class="size-4 border-gray-300 text-violet-800 focus:ring-violet-800"
                                                                       checked>
                                                            </div>
                                                            <div class="pl-7 text-sm/6">
                                                                <label for="privacy-public"
                                                                       class="font-medium text-gray-900">Public
                                                                    Access</label>
                                                                <p id="privacy-public-description"
                                                                   class="text-gray-500">
                                                                    Everyone with the link will see this product.
                                                                </p>
                                                            </div>
                                                        </div>
                                                        <div>
                                                            <div class="relative flex items-start">
                                                                <div class="absolute flex h-6 items-center">
                                                                    <input id="Public"
                                                                           name="Private"
                                                                           value="true"
                                                                           aria-describedby="privacy-private-to-project-description"
                                                                           type="radio"
                                                                           class="size-4 border-gray-300 text-violet-800 focus:ring-violet-800">
                                                                </div>
                                                                <div class="pl-7 text-sm/6">
                                                                    <label for="privacy-private-to-project"
                                                                           class="font-medium text-gray-900">Private</label>
                                                                    <p id="privacy-private-to-project-description"
                                                                       class="text-gray-500">
                                                                        Only members of this product would be able to
                                                                        access.
                                                                    </p>
                                                                </div>
                                                            </div>
                                                        </div>
                                                    </div>
                                                </fieldset>
                                            </div>

                                            ${this.submitFailed ? html`
                                                <error-alert ${ref(this.errorAlertRef)}>
                                                    <span slot="message">${this.errorData.message}</span>
                                                    <ul role="list" class="list-disc space-y-1 pl-5 mt-2"
                                                        slot="description">
                                                        ${map(this.errorData.helpTexts, (text) => html`
                                                            <li>${text}</li>`
                                                        )}
                                                    </ul>
                                                </error-alert>` : ``}
                                        </form>
                                    </div>
                                    <div class="flex shrink-0 justify-end px-4 py-4">
                                        <button type="reset"
                                                form="product-creation-form"
                                                class="rounded-md bg-white px-3 py-2 text-sm font-semibold text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 hover:bg-gray-50">
                                            Cancel
                                        </button>
                                        <button type="submit"
                                                form="product-creation-form"
                                                class="ml-4 inline-flex justify-center rounded-md bg-violet-800 px-3 py-2 text-sm font-semibold text-white shadow-sm hover:bg-violet-700 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-violet-800">
                                            Save
                                        </button>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </dialog>
        `;
    }
}

declare global {
    interface HTMLElementTagNameMap {
        'product-creation-drawer': ProductCreationDrawer;
    }
}
