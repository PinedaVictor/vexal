import { useState, useEffect } from "react";

function useArea(initialWidth: number, initialHeight: number): number {
  const [width, setWidth] = useState(initialWidth);
  const [height, setHeight] = useState(initialHeight);
  const [area, setArea] = useState(initialWidth * initialHeight);

  useEffect(() => {
    setArea(width * height);
  }, [width, height]);

  const setDimensions = (newWidth: number, newHeight: number) => {
    setWidth(newWidth);
    setHeight(newHeight);
  };

  return { width, height, area, setDimensions };
}

export default useArea;
