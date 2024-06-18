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
import {RouteItem} from "../applogic/model";


/* yankiebar */

export default function ReadRoute(props: { id: number }) {
    const { id } = props;
    const [transactionId, settransactionId] = useState(0);
    const readResult = useService<RouteItem>(
      "magic-mix.route",
      ["read", id?.toString()],
      "",
      6000,
      transactionId.toString()
    );
    const route = readResult.data;
    return (
      <div>
          
    {route && <div>
        <div>
        <div className="font-bold" >Name</div>
        <div>{route.name}</div>
    </div>    <div>
        <div className="font-bold" >Description</div>
        <div>{route.description}</div>
    </div>    <div>
        <div className="font-bold" >method</div>
        <div>{route.method}</div>
    </div>    <div>
        <div className="font-bold" >slug</div>
        <div>{route.slug}</div>
    </div>    <div>
        <div className="font-bold" >api</div>
        <div>{route.api_id}</div>
    </div>
    <div>
        <div>
        <div className="font-bold" >id</div>
        <div>{route.id}</div>
    </div>
        <div>
        <div className="font-bold" >created_at</div>
        <div>{route.created_at}</div>
    </div>
        <div>
        <div className="font-bold" >created_by</div>
        <div>{route.created_by}</div>
    </div>
        <div>
        <div className="font-bold" >updated_at</div>
        <div>{route.updated_at}</div>
    </div>
        <div>
        <div className="font-bold" >updated_by</div>
        <div>{route.updated_by}</div>
    </div>
    </div>
    </div>}


     
      </div>
    );
  }
      
