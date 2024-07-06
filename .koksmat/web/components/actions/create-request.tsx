"use client";

import Process from "@/app/koksmat/process";
import { Button } from "@/components/ui/button";
import { useState } from "react";

export default function CreateRequest() {
  const [running, setrunning] = useState(false);
  const [error, seterror] = useState("");
  const [result, setresult] = useState("");
  return (
    <div>
      <Button
        disabled={running}
        onClick={() => {
          setresult("");
          seterror("");
          setrunning(true);
        }}
      >
        Set default country
      </Button>
      {error && <div className="bg-red-500">Error: {error}</div>}
      {running && (
        <div>
          Processing:
          <Process
            servicename="magic-mix.app"
            processname="user_setcountry"
            payload={{
              user_id: 1,
              country_id: 1,
            }}
            onError={(error: any) => {
              setrunning(false);
              seterror(error);
            }}
            onSuccess={(result: any) => {
              setrunning(false);
              if (result.hasError) {
                seterror(result.errorMessage);
              } else {
                setresult(result.data);
              }
            }}
          />
        </div>
      )}
    </div>
  );
}
