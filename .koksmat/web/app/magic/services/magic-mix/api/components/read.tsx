    /* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/ 
"use client";
import { useService } from "@/app/koksmat/useservice";
import { useState } from "react";
import {APIItem} from "../applogic/model";


/* yankiebar */

export default function ReadAPI(props: { id: number }) {
    const { id } = props;
    const [transactionId, settransactionId] = useState(0);
    const readResult = useService<APIItem>(
      "magic-mix.api",
      ["read", id?.toString()],
      "",
      6000,
      transactionId.toString()
    );
    const api = readResult.data;
    return (
      <div>
          
    {api && <div>
        <div>
        <div className="font-bold" >Name</div>
        <div>{api.name}</div>
    </div>    <div>
        <div className="font-bold" >Description</div>
        <div>{api.description}</div>
    </div>    <div>
        <div className="font-bold" >method</div>
        <div>{api.method}</div>
    </div>                <div>
                    <div className="font-bold" >source</div>
                    <div>{JSON.stringify(api.source,null,2)}</div>
                </div>                <div>
                    <div className="font-bold" >schema</div>
                    <div>{JSON.stringify(api.schema,null,2)}</div>
                </div>
    <div>
        <div>
        <div className="font-bold" >id</div>
        <div>{api.id}</div>
    </div>
        <div>
        <div className="font-bold" >created_at</div>
        <div>{api.created_at}</div>
    </div>
        <div>
        <div className="font-bold" >created_by</div>
        <div>{api.created_by}</div>
    </div>
        <div>
        <div className="font-bold" >updated_at</div>
        <div>{api.updated_at}</div>
    </div>
        <div>
        <div className="font-bold" >updated_by</div>
        <div>{api.updated_by}</div>
    </div>
    </div>
    </div>}


     
      </div>
    );
  }
      
