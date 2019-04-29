
# Update all git repositories.
all-update:
	ghq list | ghq import -update

# Pull all submdules.
gsfgpo:
	git submodule foreach git pull origin
