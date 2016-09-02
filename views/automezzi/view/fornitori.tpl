            
 <div class="col-md-10 content">
    <div class="panel panel-default">
        <div class="panel-heading">
            <h3>Gestione Utenti</h3>
        </div>
        <div class="panel-body">           
            {{if .flash.error}}
            <div class="alert alert-danger">
                <button type="button" class="close" data-dismiss="alert" aria-hidden="true">
                    ×</button>
                <span class="glyphicon glyphicon-hand-right"></span> <strong> ATTENZIONE - Errore</strong>
                <hr class="message-inner-separator">
                {{.flash.error}}.
            </div>
            {{end}}
            {{if .flash.notice}}
            <div class="alert alert-success">
                <button type="button" class="close" data-dismiss="alert" aria-hidden="true">
                    ×</button>
               <span class="glyphicon glyphicon-ok"></span> <strong>Operazione completata con successo</strong>
                <hr class="message-inner-separator">
                {{.flash.notice}}.
            </div>
            {{end}}            
            {{if .Errors}}
            {{range $rec := .Errors}}
            <div class="alert alert-danger">
                <button type="button" class="close" data-dismiss="alert" aria-hidden="true">
                    ×</button>
                <span class="glyphicon glyphicon-hand-right"></span> <strong> ATTENZIONE - Errore</strong>
                <hr class="message-inner-separator">
                {{$rec}}.
            </div>      
            {{end}}
            &nbsp;
            {{end}} 
            <div class="form-group  col-lg-2">
                <div class="btn-toolbar pull-left">
                    <button class="btn btn-primary" onclick="location.href='http://{{.domainname}}/automezzi/add/fornitore'">+ Nuovo Fornitore</button>
                </div> 
            </div>            
            <form class="form-horizontal" method="POST">
                {{.xsrfdata}}
                      
                <div class="form-group col-lg-6" >
                    <select name="comparefield">
                        <option value="Descrizione" selected="selected">Descrizione</option>
                        <option value="PI">Partita Iva</option>
                        
                    </select>

                    <select name="compareop">
                        <option value="__icontains" selected="selected">Contiene</option>
                        <option value="__not__icontains">Non Contiene</option>
                    </select>

                    {{if .Errors.Compareval}}{{.Errors.Compareval}}{{end}}
                                
                    <input name="compareval" value=""  title="Tutte le ricerche sono comparate a stringhe." />
                                                    
                    <button class="btn btn-primary" type="submit" value="Cerca" />Cerca</button>

                </div>    
                       
  
            </form> 
            <div class="form-group col-lg-6" >      
                <span>Totale: {{.count}} records –  {{.query}} ordinati per {{.order}}    </span>
            </div>         

            <div class="filterable">
                    <table class="table">
                     <thead>
                        <tr class="filters">
                            <th><input type="text" id="descrsearch" class="form-control" placeholder="Descrizione" disabled ></th>
                            <th><input type="text" id="pisearch" class="form-control" placeholder="Partita Iva" disabled hidden></th>
                            <th></th>
                            <th>
                                <button class="btn btn-default btn-xs btn-filter"><span class="glyphicon glyphicon-filter"></span> Filter</button>
                            </th>
                            </tr>

                    </thead>     
                    <thead>
                        <tr class="filters">
                            <th><a href="http://{{.domainname}}/automezzi/view/fornitori/{{if eq .order "descrizione"}}-{{end}}descrizione!{{.offset}}!{{.query}}">Descrizione</a></th>
                            <th><a href="http://{{.domainname}}/automezzi/view/fornitori/{{if eq .order "pi"}}-{{end}}pi!{{.offset}}!{{.query}}">Partita IVA</a></th>
                            <th><a href="#">Modifica</a></th> 
                            <th><a href="#">Nuovo Contratto</a></th>                         
                        </tr>
                    </thead>
                    <tbody>
                        {{.Rows}}
                    </tbody>
                    </table>
                {{if .ShowNav}}
                    <br> 
                    
                    <div class="form-group col-lg-2" >      
                            <div id="progressbar"></div>pointer in data set
                            </div>  
                            <div class="form-group col-lg-8">
                                <nav>
                                    <ul class="pager">
                                        <li>
                                            <a href="http://{{.domainname}}/automezzi/view/fornitori/{{.order}}!0!{{.query}}">&lt;&lt; Inizio</a>
                                        </li>
                                        {{if .showprev}}
                                            <li>
                                                <a href="http://{{.domainname}}/automezzi/view/fornitori/{{.order}}!{{.prev}}!{{.query}}">&lt; Precendente</a>
                                            </li>
                                        {{end}}
                                        {{if .next}}
                                            <li>
                                                <a href="http://{{.domainname}}/automezzi/view/fornitori/{{.order}}!{{.next}}!{{.query}}">Successivo &gt;</a>
                                            </li>
                                        {{end}}
                                        <li>
                                            <a href="http://{{.domainname}}/automezzi/view/fornitori/{{.order}}!{{.end}}!{{.query}}">Fine &gt;&gt;</a>
                                        </li>
                                    </ul>
                                </nav>
                            </div>
                    </div>
                {{end}}
            </div>                  
        </div>
    </div>
</div>