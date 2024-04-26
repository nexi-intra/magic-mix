/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/
"use client";

import { MagicboxContext } from "@/koksmat/magicbox-context";

import serviceSearch from "@/services/magic-mix/endpoints/service/search";
import { useContext, useState } from "react";
import { Button } from "@/components/ui/button";
import { ShowCodeFragment } from "@/services/ShowCodeFragment";
import { TestServicesCall } from "@/services/testserviceexecute";

export default function TestServiceComponent() {
  const magicbox = useContext(MagicboxContext);
  const { transactionId } = magicbox;

  const [input, setInput] = useState<any>();
  const [output, setOutput] = useState<any>();
  const [errorMessage, seterrorMessage] = useState("");
  const invokeServiceEndpoint = async () => {
    setOutput(null);
    seterrorMessage("");

    const result = await serviceSearch(transactionId, input);
    if (result.hasError) {
      seterrorMessage(result.errorMessage ?? "Unknown error");
      return;
    }
    if (result.data) {
      setOutput(result.data);
    }
  };

  return (
    <div>
      <div>
        <div>
          <div className="flex"></div>
          <div>
            <div className="text-xl my-4 spy-l-2">Test</div>
            <textarea
              style={{ height: "50vh" }}
              className="w-full border border-gray-300 rounded-lg p-2 h-1/3"
              value={input}
              onChange={(e) => setInput(e.target.value)}
            />
            <div className="p-3">
              <Button onClick={invokeServiceEndpoint}>Invoke</Button>
              <div></div>
            </div>
          </div>
          <pre>
            {JSON.stringify(
              { errorMessage, input, output, transactionId },
              null,
              2
            )}
          </pre>{" "}
        </div>
      </div>
      <ShowCodeFragment
        title="Import statement"
        code={`     
import { serviceSearch } from "@/services/magic-mix/endpoints/service/search";
`}
      />
      <ShowCodeFragment
        title="Call"
        code={`     
const [input, setInput] = useState<any>();
const [output, setOutput] = useState<any>();
const [errorMessage, seterrorMessage] = useState("");

const invokeServiceEndpoint = async () => {
    setOutput(null);
    seterrorMessage("");
const result = await serviceSearch(transactionId,input);
if (result.hasError) {
    seterrorMessage(result.errorMessage ?? "Unknown error");
    return;
}
if (result.data) {
                setOutput(result.data)

}
};`}
      />
    </div>
  );
}
