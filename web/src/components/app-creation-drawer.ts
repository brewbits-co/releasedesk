import {html, LitElement} from 'lit';
import {customElement, property, state} from 'lit/decorators.js';
import {encodeFormData} from "../utils";
import {Ref, createRef, ref} from 'lit/directives/ref.js';
import {map} from 'lit/directives/map.js';

@customElement('app-creation-drawer')
export class AppCreationDrawer extends LitElement {
    override createRenderRoot() {
        return this;
    }

    @property({type: String})
    slug = "";

    @property({type: String})
    name = "";

    @property({type: Boolean})
    showDrawer = false;

    @state()
    submitFailed = false;

    errorAlertRef: Ref<HTMLDivElement> = createRef();

    @state()
    errorData: ErrorResponse = {message: '', helpTexts: []};

    handleSubmit(event: SubmitEvent) {
        event.preventDefault();
        this.submitFailed = false;

        const formData = new FormData(event.target as HTMLFormElement);

        fetch(`/internal/products/${this.slug}/apps`, {
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
                window.location.href = `/dashboard/${this.slug}/apps/${formData.get("Name")}/details`;
            }
        });
    }

    handleClose() {
        this.showDrawer = false;
        this.submitFailed = false;
    }

    override render() {
        return html`
            <div class="text-center">
                <button type="button"
                        class="inline-flex items-center rounded-md bg-violet-900 gap-x-2 px-3 py-2 text-sm font-semibold text-white shadow-sm hover:bg-violet-800 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-violet-900"
                        @click="${() => this.showDrawer = true}">
                    <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor" class="size-5">
                        <path d="M6 3a3 3 0 0 0-3 3v2.25a3 3 0 0 0 3 3h2.25a3 3 0 0 0 3-3V6a3 3 0 0 0-3-3H6ZM15.75 3a3 3 0 0 0-3 3v2.25a3 3 0 0 0 3 3H18a3 3 0 0 0 3-3V6a3 3 0 0 0-3-3h-2.25ZM6 12.75a3 3 0 0 0-3 3V18a3 3 0 0 0 3 3h2.25a3 3 0 0 0 3-3v-2.25a3 3 0 0 0-3-3H6ZM17.625 13.5a.75.75 0 0 0-1.5 0v2.625H13.5a.75.75 0 0 0 0 1.5h2.625v2.625a.75.75 0 0 0 1.5 0v-2.625h2.625a.75.75 0 0 0 0-1.5h-2.625V13.5Z"/>
                    </svg>
                    Add Application
                </button>
            </div>

            <headless-drawer ?open="${this.showDrawer}"
                             @closed="${this.handleClose}"
                             headerTitle="New Application">
                <form slot="body" id="product-creation-form" class="px-4 sm:px-6 pb-4"
                      @reset="${this.handleClose}"
                      @submit="${this.handleSubmit}">
                    <div class="space-y-6 pb-5 pt-6">
                        <div>
                            <label for="project-name" class="block text-sm/6 font-medium text-gray-900">Name</label>
                            <div class="mt-2">
                                <input type="text"
                                       name="Name"
                                       id="Name"
                                       class="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-violet-800 sm:text-sm/6"
                                       maxlength="255"
                                       autocomplete="off"
                                       value="${this.name}"
                                       required>
                            </div>
                        </div>
                        <fieldset>
                            <legend class="text-sm/6 font-medium text-gray-900">Platform</legend>
                            <div class="mt-2 space-y-4">
                                <div class="relative flex items-start">
                                    <div class="absolute flex h-6 items-center">
                                        <input id="Platform"
                                               name="Platform"
                                               value="Android"
                                               aria-describedby="privacy-public-description"
                                               type="radio"
                                               class="size-4 border-gray-300 text-violet-800 focus:ring-violet-800"
                                               checked>
                                    </div>
                                    <div class="pl-7 text-sm/6" translate="no">
                                        <label for="privacy-private-to-project" class="font-medium text-gray-900">Android</label>
                                    </div>
                                </div>
                                <div>
                                    <div class="relative flex items-start">
                                        <div class="absolute flex h-6 items-center">
                                            <input id="Platform"
                                                   name="Platform"
                                                   value="iOS"
                                                   aria-describedby="privacy-private-to-project-description"
                                                   type="radio"
                                                   class="size-4 border-gray-300 text-violet-800 focus:ring-violet-800">
                                        </div>
                                        <div class="pl-7 text-sm/6" translate="no">
                                            <label for="privacy-private-to-project" class="font-medium text-gray-900">iOS</label>
                                        </div>
                                    </div>
                                </div>
                                <div>
                                    <div class="relative flex items-start">
                                        <div class="absolute flex h-6 items-center">
                                            <input id="Platform"
                                                   name="Platform"
                                                   value="Windows"
                                                   aria-describedby="privacy-private-to-project-description"
                                                   type="radio"
                                                   class="size-4 border-gray-300 text-violet-800 focus:ring-violet-800">
                                        </div>
                                        <div class="pl-7 text-sm/6" translate="no">
                                            <label for="privacy-private-to-project" class="font-medium text-gray-900">Windows</label>
                                        </div>
                                    </div>
                                </div>
                                <div>
                                    <div class="relative flex items-start">
                                        <div class="absolute flex h-6 items-center">
                                            <input id="Platform"
                                                   name="Platform"
                                                   value="macOS"
                                                   aria-describedby="privacy-private-to-project-description"
                                                   type="radio"
                                                   class="size-4 border-gray-300 text-violet-800 focus:ring-violet-800">
                                        </div>
                                        <div class="pl-7 text-sm/6" translate="no">
                                            <label for="privacy-private-to-project" class="font-medium text-gray-900">macOS</label>
                                        </div>
                                    </div>
                                </div>
                                <div>
                                    <div class="relative flex items-start">
                                        <div class="absolute flex h-6 items-center">
                                            <input id="Platform"
                                                   name="Platform"
                                                   value="Linux"
                                                   aria-describedby="privacy-private-to-project-description"
                                                   type="radio"
                                                   class="size-4 border-gray-300 text-violet-800 focus:ring-violet-800">
                                        </div>
                                        <div class="pl-7 text-sm/6" translate="no">
                                            <label for="privacy-private-to-project" class="font-medium text-gray-900">Linux</label>
                                        </div>
                                    </div>
                                </div>
                                <div>
                                    <div class="relative flex items-start">
                                        <div class="absolute flex h-6 items-center">
                                            <input id="Platform"
                                                   name="Platform"
                                                   value="Other"
                                                   aria-describedby="privacy-private-to-project-description"
                                                   type="radio"
                                                   class="size-4 border-gray-300 text-violet-800 focus:ring-violet-800">
                                        </div>
                                        <div class="pl-7 text-sm/6">
                                            <label for="privacy-private-to-project" class="font-medium text-gray-900">Other</label>
                                        </div>
                                    </div>
                                </div>
                            </div>
                        </fieldset>
                    </div>

                    ${this.submitFailed ? html`
                        <error-alert ${ref(this.errorAlertRef)}>
                            <span slot="message">${this.errorData.message}</span>
                            <ul role="list" class="list-disc space-y-1 pl-5 mt-2" slot="description">
                                ${map(this.errorData.helpTexts, (text) => html`
                                    <li>${text}</li>`
                                )}
                            </ul>
                        </error-alert>` : ``}
                </form>

                <span slot="footer">
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
        </span>
            </headless-drawer>
        `;
    }
}

declare global {
    interface HTMLElementTagNameMap {
        'app-creation-drawer': AppCreationDrawer;
    }
}
