import { useState } from "react";

const useLoading = () => {
  const [isLoading, setLoading] = useState(false);

  const setLoadingState = (loadingState) => {
    setLoading(loadingState);
  };

  return [isLoading, setLoadingState];
};

export default useLoading;
