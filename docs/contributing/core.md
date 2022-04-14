# Core Contribution

To make a contribution to the nuclei core, make sure you have Nuclei installed from code. If not already, you can use the below git command to install a local development copy of nuclei from source.

- Install nuclei from source

```
git clone https://github.com/projectdiscovery/nuclei.git; \
cd nuclei/v2/cmd/nuclei; \
go build; \
./nuclei -version;
```

- The next step is to understand the code structure of the Nuclei Project. [Code Structure](../references/code-structure.md) can be read to get an idea. Other useful guides are [Nuclei Go Usage Example](../guides/code/nuclei-go-example.md) and [New Nuclei Protocol Addition Example](../guides/code/new-nuclei-protocol.md). These can be used as a starting point to get a feel for the codebase.

- Now we can make our change to the code. Nuclei Core Pull Requests are made to the **dev** branch. To switch your local copy to the dev branch, type the following command.

```
git checkout dev

# To confirm, should say "dev"
git branch --show-current
```

- After this is done, all remains is making your change and then submitting a PR upstream. The code will be reviewed and after meeting the requirements, it will be merged!