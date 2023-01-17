export const lazy = (image: HTMLImageElement, src: string) => {
  const loaded = () => {
    image.style.opacity = '1';
  };

  const observer = new IntersectionObserver((entries) => {
    if (entries[0].isIntersecting) {
      image.src = src;
      if (image.complete) {
        loaded();
      } else {
        image.addEventListener('load', loaded);
      }
    }
  });

  observer.observe(image);

  return {
    destroy() {
      image.removeEventListener('load', loaded);
    }
  };
};