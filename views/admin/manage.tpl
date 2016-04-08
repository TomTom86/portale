            
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
                    <button class="btn btn-primary" onclick="location.href='http://{{.domainname}}/admin/add/id!0!id__gte,0'">+ Nuovo Utente</button>
                </div> 
            </div>            
            <form class="form-horizontal" method="POST">
                {{.xsrfdata}}
                      
                <div class="form-group col-lg-6" >
                    <select name="comparefield">
                        <option value="id" selected="selected">Id</option>
                        <option value="first">First</option>
                        <option value="last">Last</option>
                        <option value="email">Email</option>
                        <option value="id_key">ID key</option>
                        <option value="reg_date">Reg date</option>
                    </select>

                    <select name="compareop">
                        <option value="__exact">=</option>
                        <option value="__not__exact">!=</option>
                        <option value="__lt">&lt;</option>
                        <option value="__lte">&lt;=</option>
                        <option value="__gt">&gt;</option>
                        <option value="__gte" selected="selected">&gt;=</option>
                        <option value="__contains">contains</option>
                        <option value="__not__contains">not contains</option>
                        <option value="__icontains">icontains</option>
                        <option value="__not__icontains">not icontains</option>
                    </select>

                    {{if .Errors.Compareval}}{{.Errors.Compareval}}{{end}}
                                
                    <input name="compareval" value="0"  title="All search terms are compared as text strings. 'Reg date' has a date format of yyyy-mm-dd." />
                                                    
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
                            <th><input type="text" id="idsearch" class="form-control" placeholder="ID" disabled ></th>
                            <th><input type="text" id="firstsearch"  class="form-control" placeholder="Nome" disabled ></th>
                            <th><input type="text" id="lastsearch"  class="form-control" placeholder="Cognome" disabled hidden></th>
                            <th><input type="text" id="emailsearch" class="form-control" placeholder="Email" disabled hidden></th>
                            <th>
                                <div class="pull-right">
                                    <button class="btn btn-default btn-xs btn-filter"><span class="glyphicon glyphicon-filter"></span> Filter</button>
                                </div>
                            </th>
                            </tr>

                    </thead>   
                    <thead>
                        <tr class="filters">
                            <th><a href="http://{{.domainname}}/admin/{{if eq .order "id"}}-{{end}}id!{{.offset}}!{{.query}}">ID</a></th>
                            <th><a href="http://{{.domainname}}/admin/{{if eq .order "first"}}-{{end}}first!{{.offset}}!{{.query}}">Nome</a></th>
                            <th><a href="http://{{.domainname}}/admin/{{if eq .order "last"}}-{{end}}last!{{.offset}}!{{.query}}">Cognome</a></th>
                            <th><a href="http://{{.domainname}}/admin/{{if eq .order "email"}}-{{end}}email!{{.offset}}!{{.query}}">Email</a></th>
                            <th><a href="#">Modifica</a></th>                        
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
                                            <a href="http://{{.domainname}}/admin/{{.order}}!0!{{.query}}">&lt;&lt; Inizio</a>
                                        </li>
                                        {{if .showprev}}
                                            <li>
                                                <a href="http://{{.domainname}}/admin/{{.order}}!{{.prev}}!{{.query}}">&lt; Precendente</a>
                                            </li>
                                        {{end}}
                                        {{if .next}}
                                            <li>
                                                <a href="http://{{.domainname}}/admin/{{.order}}!{{.next}}!{{.query}}">Successivo &gt;</a>
                                            </li>
                                        {{end}}
                                        <li>
                                            <a href="http://{{.domainname}}/admin/{{.order}}!{{.end}}!{{.query}}">Fine &gt;&gt;</a>
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