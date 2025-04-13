/**
 * Encodes a FormData object into a URL-encoded string.
 *
 * The function iterates through each entry in the FormData object,
 * appending each key-value pair to a URLSearchParams instance.
 * It then returns the resulting URL-encoded string.
 *
 * @param {FormData} data - The FormData object to be encoded.
 * @returns {string} A URL-encoded string representing the key-value pairs in the given FormData object.
 */
export const encodeFormData = (data: FormData): string => {
  const params = new URLSearchParams();
  for (const [key, value] of data.entries()) {
    params.append(key, value as string);
  }
  return params.toString();
}

/**
 * Converts a given string into a URL-friendly slug.
 *
 * A slug is a lowercase string with spaces replaced by hyphens and
 * non-alphanumeric characters removed. Useful for creating clean, readable URLs.
 *
 * @param {string} name - The input string to convert into a slug.
 * @returns {string} - The URL-friendly slug.
 *
 * @example
 * convertToSlug("Hello, World! 2024");
 * // Returns: "hello-world-2024"
 */
export const convertToSlug = (name: string): string => {
  // Convert the name to lowercase
  name = name.toLowerCase();

  // Replace non-alphanumeric characters with an empty string
  // Allow spaces as they will be replaced later
  name = name.replace(/[^a-z0-9 -]/g, "");

  // Replace spaces with hyphens
  return name.replace(/\s+/g, "-");
}