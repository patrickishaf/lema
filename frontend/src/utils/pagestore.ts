export const pageStore = {
  saveUserPage(page: number) {
    sessionStorage.setItem('upg', `${page}`);
  },

  getUserPage(): number {
    const value = sessionStorage.getItem('upg');
    if (!value) return 1
    return parseInt(value);
  },

  clearPostsPage() {
    sessionStorage.removeItem('ppg');
  },

  clear() {
    sessionStorage.removeItem('upg');
  }
}