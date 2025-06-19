import {html, LitElement} from 'lit';
import {customElement, property, state} from 'lit/decorators.js';
import {encodeFormData} from "../utils";
import {Ref, createRef, ref} from 'lit/directives/ref.js';
import {map} from 'lit/directives/map.js';

interface Channel {
  ID: number;
  Name: string;
}

@customElement('release-update-form')
export class ReleaseUpdateForm extends LitElement {
  override createRenderRoot() {
    return this;
  }

  @property({type: Number})
  releaseID = 0;

  @property({type: Number})
  targetChannel = 0;

  @property({type: String})
  status = '';

  @property({type: String})
  buildSelection = 'Last';

  @property({type: Array})
  channels: Channel[] = [];

  @state()
  submitFailed = false;

  errorAlertRef: Ref<HTMLDivElement> = createRef();

  @state()
  errorData: {message: string, helpTexts: string[]} = {message: '', helpTexts: []};

  handleBuildSelectionChange(event: Event) {
    const target = event.target as HTMLInputElement;
    if (target.checked) {
      this.buildSelection = target.value;
    }
  }

  handleSubmit(event: SubmitEvent) {
    event.preventDefault();
    this.submitFailed = false;

    const formData = new FormData(event.target as HTMLFormElement);

    fetch(`/internal/releases/${this.releaseID}`, {
      method: 'PUT',
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
        // Redirect back to the release summary page
        window.location.reload();
      }
    });
  }

  override render() {
    return html`
      <form id="release-update-form" @submit="${this.handleSubmit}">
        <div class="sm:col-span-3">
          <label for="TargetChannel" class="block text-sm/6 font-medium text-violet-900 mb-2">Channel</label>
          <select id="TargetChannel" name="TargetChannel" autocomplete="off"
                  class="col-start-1 row-start-1 w-full appearance-none rounded-md bg-white py-1.5 pl-3 pr-8 text-base text-gray-900 outline outline-1 -outline-offset-1 outline-gray-300 focus:outline focus:outline-2 focus:-outline-offset-2 focus:outline-violet-600 sm:text-sm/6">
            ${this.channels.map(channel => html`
              <option value="${channel.ID}" ?selected="${channel.ID === this.targetChannel}">${channel.Name}</option>
            `)}
          </select>
        </div>

        <div class="sm:col-span-3 mt-4">
          <label for="Status" class="block text-sm/6 font-medium text-violet-900 mb-2">Status</label>
          <select id="Status" name="Status" autocomplete="off"
                  class="col-start-1 row-start-1 w-full appearance-none rounded-md bg-white py-1.5 pl-3 pr-8 text-base text-gray-900 outline outline-1 -outline-offset-1 outline-gray-300 focus:outline focus:outline-2 focus:-outline-offset-2 focus:outline-violet-600 sm:text-sm/6">
            <option value="Draft" ?selected="${this.status === 'Draft'}">Draft</option>
            <option value="Published" ?selected="${this.status === 'Published'}">Published</option>
            <option value="Deprecated" ?selected="${this.status === 'Deprecated'}">Deprecated</option>
            <option value="Unpublished" ?selected="${this.status === 'Unpublished'}">Unpublished</option>
            <option value="Scheduled" ?selected="${this.status === 'Scheduled'}">Scheduled</option>
          </select>
        </div>

        <div class="col-span-full mt-4">
          <label class="block text-sm/6 font-medium text-violet-900 mb-2">Build Selection</label>

          <fieldset>
            <div class="grid grid-cols-1 gap-4 sm:grid-cols-2">
              ${radioOption('BuildSelection', 'Last', 'Last', 'Always use the latest build.', this.buildSelection === 'Last', this.handleBuildSelectionChange)}
              <span class="pointer-events-none opacity-50 select-none">
                ${radioOption('BuildSelection', 'Manual', 'Manual', 'Select builds manually.', this.buildSelection === 'Manual', this.handleBuildSelectionChange)}
              </span>
            </div>
          </fieldset>
        </div>
          
        ${this.buildSelection === 'Manual' ? html`
          <div class="col-span-full mt-4">
            <div class="mt-2 flex justify-center rounded-lg border border-dashed border-gray-900/25 px-6 py-10">
              
            </div>
          </div>
        ` : ''}

        <div class="mt-6 flex items-center justify-end gap-x-6">
          <button type="submit" class="rounded-md bg-violet-600 px-3 py-2 text-sm font-semibold text-white shadow-sm hover:bg-violet-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-violet-600">
            Save
          </button>
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
    `;
  }
}

const radioOption = (name: string, value: string, title: string, subtitle: string, checked: boolean = false, handler?: any) => html`
  <label class="relative flex cursor-pointer rounded-lg border bg-white p-4 shadow-sm focus:outline-none">
    <input type="radio" name="${name}" value="${value}" class="peer sr-only" @change="${handler}" required
           ?checked="${checked}">
        <span class="flex flex-1">
            <span class="flex flex-col">
                <span class="block text-sm font-medium text-gray-900">${title}</span>
                <span class="mt-1 flex items-center text-sm text-gray-500">${subtitle}</span>
            </span>
        </span>

        <svg class="size-5 text-violet-800 hidden peer-checked:block" viewBox="0 0 20 20" fill="currentColor"
             aria-hidden="true" data-slot="icon">
            <path fill-rule="evenodd"
                  d="M10 18a8 8 0 1 0 0-16 8 8 0 0 0 0 16Zm3.857-9.809a.75.75 0 0 0-1.214-.882l-3.483 4.79-1.88-1.88a.75.75 0 1 0-1.06 1.061l2.5 2.5a.75.75 0 0 0 1.137-.089l4-5.5Z"
                  clip-rule="evenodd"/>
        </svg>
        <span class="pointer-events-none absolute -inset-px rounded-lg border-2 peer-checked:border-violet-800"
              aria-hidden="true"></span>
    </label>
`

declare global {
  interface HTMLElementTagNameMap {
    'release-update-form': ReleaseUpdateForm;
  }
}
