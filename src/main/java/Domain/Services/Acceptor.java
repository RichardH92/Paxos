package Domain.Services;

import Domain.Messages.Accept;
import Domain.Messages.Prepare;
import Domain.Messages.Promise;
import Domain.Messages.Propose;

public class Acceptor {
    private Propose highestProposal;

    public Promise handlePrepare(Prepare prepareMsg) {
        return null;
    }

    public Accept handlePropose(Propose proposal) {
        return null;
    }

}
