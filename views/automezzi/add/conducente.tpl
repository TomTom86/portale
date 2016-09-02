

<div class="col-md-10 content">
    <div class="panel panel-default">
        <div class="panel-heading">
            <h3>Inserimento Conducente</h3>
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
                <form class="form-horizontal" method="POST">
                    {{ .xsrfdata }}         
                    <!-- NOME --> 
                    <div class="form-group">                
                        <label class="col-md-4 control-label" for="nome">Nome</label>                          
                        <div class="col-md-4">
                        <input id="condnome" name="condnome" type="text" value="{{.CondNome}}" class="form-control input-md"/>
                        </div> 
                    </div> 
                    <!-- COGNOME --> 
                    <div class="form-group">
                        <label class="col-md-4 control-label" for="cognome">Cognome</label>                     
                        <div class="col-md-4">
                        <input id="condcognome" name="condcognome" type="text" value="{{.CondCognome}}" class="form-control input-md"/>
                        </div> 
                    </div>
                    <!-- CODICE FISCALE -->            
                    <div class="form-group">
                        <label class="col-md-4 control-label" for="cf">Codice Fiscale</label>
                        <div class="col-md-4">
                        <input id="condcf" name="condcf" type="text" value="{{.CondCF}}" class="form-control input-md"/>
                        </div>
                    </div>  
                    
                    <!-- footer --> 
                    <div class="col-md-12 form-group">
                        <button class="btn btn-primary pull-right" type="submit" class="pull-right" value="Inserisci" />Inserisci</button>
                        
                                  
                    </div>                              

                </form>
        </div>
    </div>
</div>