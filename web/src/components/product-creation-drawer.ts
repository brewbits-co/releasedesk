import {html, LitElement} from 'lit';
import {customElement, property, state} from 'lit/decorators.js';
import {convertToSlug, encodeFormData} from "../utils";
import {Ref, createRef, ref} from 'lit/directives/ref.js';
import {map} from 'lit/directives/map.js';

@customElement('product-creation-drawer')
export class ProductCreationDrawer extends LitElement {
  override createRenderRoot() {
    return this;
  }

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

  handleClose() {
    this.showDrawer = false;
    this.submitFailed = false;
  }

  override render() {
    return html`
      <div class="text-center">
        <button type="button"
                class="inline-flex items-center rounded-md bg-violet-900 px-3 py-2 text-sm font-semibold text-white shadow-sm hover:bg-violet-800 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-violet-900"
                @click="${() => this.showDrawer = true}">
          <i class="fa-solid fa-layer-group"></i>&nbsp;&nbsp;Create Your Product
        </button>
      </div>

      <headless-drawer ?open="${this.showDrawer}"
                       @closed="${this.handleClose}"
                       headerTitle="New Product"
                       headerSubtitle="Get started by filling in the information below to create your new product.">
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
                       required>
              </div>
            </div>
            <div>
              <label for="project-name" class="block text-sm/6 font-medium text-gray-900">Slug</label>
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
              <legend class="text-sm/6 font-medium text-gray-900">Visibility</legend>
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
                    <label for="privacy-public" class="font-medium text-gray-900">Public
                      Access</label>
                    <p id="privacy-public-description" class="text-gray-500">
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
                      <label for="privacy-private-to-project" class="font-medium text-gray-900">Private</label>
                      <p id="privacy-private-to-project-description" class="text-gray-500">
                        Only members of this product would be able to access.
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
    'product-creation-drawer': ProductCreationDrawer;
  }
}
