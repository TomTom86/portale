

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
                    <!-- DESCRIZIONE --> 
                    <div class="form-group">                
                        <label class="col-md-4 control-label" for="nome">Descrizione</label>                          
                        <div class="col-md-4">
                        <input id="fornitore" name="fornitore" type="text" value="{{.Fornitore}}" class="form-control input-md"/>
                        </div> 
                    </div> 
                    <!-- PARTITA IVA -->            
                    <div class="form-group">
                        <label class="col-md-4 control-label" for="cf">Partita Iva</label>
                        <div class="col-md-4">
                        <input id="pi" name="pi" type="text" value="{{.PI}}" class="form-control input-md"/>
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